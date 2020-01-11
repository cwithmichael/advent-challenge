(require '[clojure.string :as str])

(defn get-coords [file]
  (vec (map #(str/split % #",") (str/split (slurp file) #"\n") )))

(defn get-manhattan-distance [p q]
  ; Find the manhattan distance between two points
  (let [[px py] p
        [qx qy] q
        distance (+ (Math/abs (- qx px)) (Math/abs (- qy py)))]
    distance))

(defn get-orientation [p q r]
  ; Find orientation of ordered triplet
  (let [[px py] p
        [qx qy] q
        [rx ry] r
        orientation (- (* (- qy py) (- rx qx)) (* (- qx px) (- ry qy)))]
    (cond 
      (= orientation 0) 0
      (> orientation 0) 1
      (< orientation 0) 2)))

(defn on-segment? [p q r]
  ;Does point p lie on line segment pr?
  (let [[px py] p
        [qx qy] q
        [rx ry] r]
    (cond
      (and 
       (<= qx (max px rx))
       (>= qx (min px rx))
       (<= qy (max py ry))
       (>= qy (min py ry))) true
      :else false)))

(defn calculate-intersection-point [p q r s]
  ; Calculate the intersection point
  ; http://www.ambrsoft.com/MathCalc/Line/TwoLinesIntersection/TwoLinesIntersection.htm
  (println p q r s)
  (let [[x1 y1] p
        [x2 y2] q
        [x3 y3] r 
        [x4 y4] s
        denom (- (* (- x2 x1) (- y4 y3)) (* (- x4 x3) (- y2 y1)))
        x (/ 
           (- (* (- (* x2 y1) (* x1 y2)) (- x4 x3)) (* (- (* x4 y3) (* x3 y4)) (- x2 x1))) denom)
        y (/ 
           (- (* (- (* x2 y1) (* x1 y2)) (- y4 y3)) (* (- (* x4 y3) (* x3 y4)) (- y2 y1))) denom)]
    [x y]))

(defn intersect? [pq pq2]
  ; Do these two lines intersect?
  ;(println pq pq2)
  (let [
        [p q] pq
        [p2 q2] pq2
        orientation-1 (get-orientation p q p2) 
        orientation-2 (get-orientation p q q2) 
        orientation-3 (get-orientation p2 q2 p) 
        orientation-4 (get-orientation p2 q2 q)]
    (cond
      ; General case
      (and (not= orientation-1 orientation-2) (not= orientation-3 orientation-4)) [pq pq2]
      ; Special cases
      (and (= orientation-1 0) (on-segment? p p2 q)) [pq pq2]
      (and (= orientation-2 0) (on-segment? p q2 q)) [pq pq2]
      (and (= orientation-3 0) (on-segment? p2 p q2)) [pq pq2]
      (and (= orientation-4 0) (on-segment? p2 q q2)) [pq pq2]
      :else nil)))
    

(defn get-intersection-points [line-segments]
  ; Return a list of intersection points
  (let [x (get line-segments 0) y  (get line-segments 1) cmp (compare (count x) (count y)) intersection-points []]
    (cond 
      (empty? x) intersection-points
      (= cmp 0) (for [i (range (count x)) j (range (count x))] (conj intersection-points 
                                                           (intersect? 
                                                            (nth x i) 
                                                            (nth y j))))
      (= cmp 1) (for [i (range (count x)) j (range (count y))] (conj intersection-points 
                                                                     (intersect? 
                                                                      (nth x i) 
                                                                      (nth y j)))))))

(defn get-points [source movement line]
  ; Get points for line
  ;(println line)
  (if (empty? movement) (conj line source)
      (let [
            [_ direction steps] (re-matches #"([RLUD])(\d+)" (first movement))
            num-steps (Integer/parseInt steps)]
        (cond
          (= direction "R") (get-points
                             (assoc source 0 (+ (get source 0) num-steps)) (drop 1 movement) (conj line source))
          (= direction "U") (get-points
                             (assoc source 1 (+ (get source 1) num-steps)) (drop 1 movement) (conj line source))
          (= direction "L") (get-points
                             (assoc source 0 (- (get source 0) num-steps)) (drop 1 movement) (conj line source))
          (= direction "D") (get-points
                             (assoc source 1 (- (get source 1) num-steps)) (drop 1 movement) (conj line source))))))

(defn get-line-segments [lines]
  ; break up the lines into segments
  ;(println lines)
  (let [segments [(partition 2 1 (get lines 0)) (partition 2 1 (get lines 1))]] segments))

(defn get-lines [coords]
  ; Get lines based on coords provided in input
  (loop [i 0 lines []]
    (if (= i 2)
      lines
      (recur (inc i) (conj lines (get-points [0 0] (get coords i) []))))))

(defn find-closest [intersection-points]
  ;Find closet intersection point to origin; We don't care about points that include [0 0]
  ;(println intersection-points)
  (let [filtered-points (remove (fn [x] (or (nil? (first x)) (= (nth (get (get x 0) 0) 0) [0 0]))) intersection-points)]
    (println filtered-points)
    (loop [i 0
           cnt (count filtered-points)
           distances []]
      (if (= i cnt) 
        distances
        (recur 
         (inc i) cnt 
         (conj distances 
               (get-manhattan-distance [0 0] (calculate-intersection-point (nth (get (get (nth filtered-points i) 0) 0) 0)
                                                                           (nth (get (get (nth filtered-points i) 0) 0) 1)
                                                                           (nth (get (get (nth filtered-points i) 0) 1) 0)
                                                                           (nth (get (get (nth filtered-points i) 0) 1) 1)))))))))

(println (apply min (find-closest (get-intersection-points (get-line-segments (get-lines (get-coords "input")))))))