package face_finder

import (
	"io/ioutil"
	"mime/multipart"

	pigo "github.com/esimov/pigo/core"
)

func IsPhotoWithFace(file multipart.File) (bool, error) {
	cascadeFile, err := ioutil.ReadFile("facefinder")
	if err != nil {
		return false, err
	}
	photo, err := pigo.DecodeImage(file)
	if err != nil {
		return false, err
	}

	pixels := pigo.RgbToGrayscale(photo)
	cols, rows := photo.Bounds().Max.X, photo.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	pigoNew := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigoNew.Unpack(cascadeFile)
	if err != nil {
		return false, err
	}

	angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	dets := classifier.RunCascade(cParams, angle)

	// Calculate the intersection over union (IoU) of two clusters.
	dets = classifier.ClusterDetections(dets, 0.2)
	for _, d := range dets {
		if d.Q >= 5 {
			return true, nil
		}
	}
	return false, nil
}
