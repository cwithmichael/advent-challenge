(require '[clojure.java.io :as io])

(defn calculate-fuel [x]
 (loop [tmp x acc (- x)]
   (if (<= tmp 0)
     acc
     (recur (- (quot tmp 3) 2) (+ tmp acc)))))

(with-open [rdr (io/reader "puzzle1_input.txt")]
  (let [masses (map #(Integer/parseInt %) (reduce conj [] (line-seq rdr)))]
    (println (reduce + 0 (map calculate-fuel masses)))))
