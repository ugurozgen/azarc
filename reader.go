package main

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"runtime"
	"sync"
)

// NewOmdbTsvReader creates OmdbTsvReader with default values.
// If you provide any option it will apply options and override
// default values.
func NewOmdbTsvReader(opts ...TsvReaderOption) *OmdbTsvReader {
	ctx, cancel := context.WithCancel(context.Background())

	var (
		defaultGoroutineCount = runtime.NumCPU()
		defaultFileName       = "title.basics_test.tsv"
		defaultContext        = ctx
		defaultCancel         = cancel
	)

	otr := &OmdbTsvReader{
		goroutineCount: defaultGoroutineCount,
		fileName:       defaultFileName,
		ctx:            defaultContext,
		cancel:         defaultCancel,
		wg:             sync.WaitGroup{},
		outputCh:       make(chan OmdbTitleRecord),
		recordCh:       make(chan []string),
	}

	for _, opt := range opts {
		opt(otr)
	}

	return otr
}

// ReadAsync reads and parses each line of tsv file concurrently
func (otr *OmdbTsvReader) ReadAsync() (chan OmdbTitleRecord, error) {
	reader, err := otr.reader()
	if err != nil {
		return nil, err
	}

	otr.read(reader)

	otr.startParsers()

	otr.wait()

	return otr.outputCh, nil
}

// read reads tsv file and puts record to recordCh
func (otr *OmdbTsvReader) read(reader *csv.Reader) {
	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}

			// ignore broken lines
			if err != nil {
				continue
			}
			otr.recordCh <- record
		}
		close(otr.recordCh)
	}()
}

// startParsers runs parser goroutines to consume tsv records and map them to OmdbTitleRecord
// startParsers method works with goroutineCount
func (otr *OmdbTsvReader) startParsers() {
	for w := 0; w < otr.goroutineCount; w++ {
		otr.wg.Add(1)
		go func() {
			defer otr.wg.Done()
			otr.parserRoutine(otr.recordCh, otr.outputCh)
		}()
	}
}

// wait waits for all tsv records get done
func (otr *OmdbTsvReader) wait() {
	go func() {
		otr.wg.Wait()
		close(otr.outputCh)
	}()
}

// parserRoutine reads records from recordCh
// after parsing the record, puts OmdbTitleRecord to outputCh
func (otr *OmdbTsvReader) parserRoutine(recordCh <-chan []string, outputCh chan<- OmdbTitleRecord) {
	for {
		select {
		case record, ok := <-recordCh:
			if !ok {
				return
			}
			outputCh <- parseRecord(record)

		case <-otr.ctx.Done():
			return
		}
	}
}

// parseRecord maps a tsv record to OmdbTitleRecord
func parseRecord(record []string) OmdbTitleRecord {
	return OmdbTitleRecord{
		Tconst:         record[0],
		TitleType:      record[1],
		PrimaryTitle:   record[2],
		OriginalTitle:  record[3],
		IsAdult:        record[4],
		StartYear:      record[5],
		EndYear:        record[6],
		RuntimeMinutes: record[7],
		Genres:         record[8],
	}
}

// reader generates tsv reader
func (otr *OmdbTsvReader) reader() (*csv.Reader, error) {
	file, err := os.Open(otr.fileName)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'

	return csvReader, nil
}

// TsvReaderOption is a variadic function that is being used to apply options to TsvReader
type TsvReaderOption func(*OmdbTsvReader)

// WithGoroutineCount is a TsvReaderOption and enables setting goroutineCount for TsvReader
func WithGoroutineCount(goroutineCount int) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.goroutineCount = goroutineCount
	}
}

// Filename is a TsvReaderOption and enables setting file name for TsvReader
func Filename(name string) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.fileName = name
	}
}

// WithContext is a TsvReaderOption and enables setting context for TsvReader
func WithContext(ctx context.Context) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.ctx = ctx
	}
}

// WithCancel is a TsvReaderOption and enables setting cancel function for TsvReader
func WithCancel(cancel context.CancelFunc) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.cancel = cancel
	}
}
