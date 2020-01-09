(require '[clojure.string :as str])

(defn get-lines [file]
  (str/split (str/trim (slurp file)) #","))

(defn get-operands [ops idx]
  (vector (get ops (get ops (+ idx 1))) (get ops (get ops (+ idx 2)))))

(defn execute-op [op ops idx]
  (cond
    (= op :add) (reduce + (get-operands ops idx))
    (= op :mul) (reduce * (get-operands ops idx))))

(defn update-noun-verb [ops noun verb]
  (assoc (assoc ops 1 noun) 2 verb))

(defn opcode-parser [ops idx target]
  (let [x (get ops idx)]
    (cond
      (= (get ops 0) target) ops
      (= x 99) :done
      (= x 1) (recur (assoc ops (get ops (+ idx 3)) (execute-op :add ops idx)) (+ idx 4) target)
      (= x 2) (recur (assoc ops (get ops (+ idx 3)) (execute-op :mul ops idx)) (+ idx 4) target)
      :else (recur ops (+ idx 1) target))))

(defn opcode-parser-redux [ops target]
  (for [noun (range 99)
        verb (range 99)
        :let [ops (update-noun-verb ops noun verb) result (opcode-parser ops 0 target)]
        :when (not= result :done)]
    (+ (* 100 noun) verb)))

(defn compute [filename target]
  (let [ops (vec (map #(Integer/parseInt %) (get-lines filename)))]
    (opcode-parser-redux ops target)))

(println (compute "original_input.txt" 19690720))