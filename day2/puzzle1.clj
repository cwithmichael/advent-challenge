(require '[clojure.string :as str])

(defn get-lines [file]
  (str/split (str/trim (slurp file)) #","))

(defn get-operands [ops idx]
  (vector (get ops (get ops (+ idx 1))) (get ops (get ops (+ idx 2)))))

(defn execute-op [op ops idx]
  (cond
    (= op :add) (reduce + (get-operands ops idx))
    (= op :mul) (reduce * (get-operands ops idx))))

(defn opcode-parser [ops idx]
  (let [x (get ops idx)]
    (cond
      (= x 99) ops
      (= x 1) (recur (assoc ops (get ops (+ idx 3)) (execute-op :add ops idx)) (+ idx 4))
      (= x 2) (recur (assoc ops (get ops (+ idx 3)) (execute-op :mul ops idx)) (+ idx 4))
      :else (recur ops (+ idx 1)))))

(defn compute [filename]
  (let [ops (vec (map #(Integer/parseInt %)(get-lines filename)))]
    (opcode-parser ops 0)))

(println (compute "input"))