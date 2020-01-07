(require '[clojure.java.io :as io])

(defn calculate-fuel-sum [xs] (reduce #(+ (- (quot %2 3) 2) %1) 0 xs))

(with-open [rdr (io/reader "puzzle1_input.txt")]
  (let [masses (map #(Integer/parseInt %) (reduce conj [] (line-seq rdr)))]
    (println (calculate-fuel-sum masses))))
