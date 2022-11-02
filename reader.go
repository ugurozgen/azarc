package main

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"runtime"
	"sync"
)

type TsvReader interface {
	ReadAsync() error
}

type OmdbTsvReader struct {
	goroutineCount int
	fileName       string
	ctx            context.Context
	cancel         context.CancelFunc
	outputCh       chan OmdbTitleRecord
	wg             sync.WaitGroup
	recordCh       chan []string
}

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

func (otr *OmdbTsvReader) startParsers() {
	for w := 0; w < otr.goroutineCount; w++ {
		otr.wg.Add(1)
		go func() {
			defer otr.wg.Done()
			otr.parserRoutine(otr.recordCh, otr.outputCh)
		}()
	}
}

func (otr *OmdbTsvReader) wait() {
	go func() {
		otr.wg.Wait()
		close(otr.outputCh)
	}()
}

func (otr *OmdbTsvReader) parserRoutine(recordCh <-chan []string, titleCh chan<- OmdbTitleRecord) {
	for {
		select {
		case record, ok := <-recordCh:
			if !ok {
				return
			}
			titleCh <- parseRecord(record)

		case <-otr.ctx.Done():
			return
		}
	}
}

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

func (otr *OmdbTsvReader) reader() (*csv.Reader, error) {
	file, err := os.Open(otr.fileName)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'

	return csvReader, nil
}

type TsvReaderOption func(*OmdbTsvReader)

func WithGoroutineCount(goroutineCount int) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.goroutineCount = goroutineCount
	}
}

func Filename(name string) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.fileName = name
	}
}

func WithContext(ctx context.Context) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.ctx = ctx
	}
}

func WithCancel(cancel context.CancelFunc) TsvReaderOption {
	return func(otr *OmdbTsvReader) {
		otr.cancel = cancel
	}
}
