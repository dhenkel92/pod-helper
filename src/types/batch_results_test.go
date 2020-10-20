package types

import "testing"

func TestShouldCreateTwoEqualBatches(t *testing.T) {
	results := []Result{{}, {}, {}, {}, {}, {}}
	batches := BatchResults(3, results)
	if len(batches) != 2 {
		t.Fail()
	}
	for _, batch := range batches {
		if len(batch) != 3 {
			t.Fail()
		}
	}
}

func TestShouldCreateOnlyOneBatch(t *testing.T) {
	results := []Result{{}}
	batches := BatchResults(3, results)
	if len(batches) != 1 {
		t.Fail()
	}
	for _, batch := range batches {
		if len(batch) != 1 {
			t.Fail()
		}
	}
}

func TestShouldCreateTwoBatchesWithRest(t *testing.T) {
	results := []Result{{}, {}, {}, {}, {}, {}, {}, {}}
	batches := BatchResults(3, results)
	if len(batches) != 3 {
		t.Fail()
	}
	if len(batches[0]) != 3 {
		t.Fail()
	}
	if len(batches[1]) != 3 {
		t.Fail()
	}
	if len(batches[2]) != 2 {
		t.Fail()
	}
}

func TestShouldReturnEmptyResultWithEmptyInput(t *testing.T) {
	batches := BatchResults(3, []Result{})
	if len(batches) != 0 {
		t.Fail()
	}
}
