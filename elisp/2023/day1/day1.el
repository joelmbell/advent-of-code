(defun load-file (filename)
  "return the contents of FILENAME."
  (s-split "\n" (f-read filename) t))

(defun number? (char)
  (numberp (read (string char))))

(defun chars-to-num (chars)
  (string-to-number
   (-reduce-from
	#'concat
	""
	(-map #'string `(,(-first-item chars) ,(-last-item chars))))))

(defun solve-pt-1 (filename)
  (-reduce
   #'+
   (-map
	(lambda (row)
	  (chars-to-num
	   (-filter #'number? (string-to-list row))))
	(load-file filename))))

(solve-pt-1 "input.txt")

