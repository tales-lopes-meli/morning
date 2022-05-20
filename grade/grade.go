package grade

import "errors"

func Average (grades ...float64) (float64, error) {
	var sum float64 = 0
	var err error = nil
	var average float64 = 0

	for _, grade := range grades {
		if grade < 0 {
			err = errors.New("Negative grades are not allowed")
		} else {
			sum += grade
		}
	}
	if(err == nil) {
		average = sum / float64(len(grades))
	}
	return average, err
}