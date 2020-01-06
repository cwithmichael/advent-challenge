#!/usr/bin/env stack
import System.IO

main :: IO ()
main = do
    puzzleInput <- readFile "puzzle1_input.txt"
    putStrLn $ show $ calculateFuelSum $ map (read::String->Int) $ lines puzzleInput

calculateFuelSum :: [Int] -> Int
calculateFuelSum xs = foldr (\x y -> (x `div` 3 - 2) + y) 0 xs

