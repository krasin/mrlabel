mrLabel - simple web frontend to label images.

0) Check mrlabel.html for label names. Labels are integers, 0 is no label, 1..N is meaningful labels. N should be less than 9 for the numeric keys to work.

1) go run logger.go; this will start a logging server that will write labels to labels.txt

2) Go to http://localhost:10000/

3) for each image, press 0..N to label it. Alternatively, click on the coresponing button.

4) collect your labels from labels.txt


