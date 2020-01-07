defmodule Puzzle1 do
    def calculate_fuel_sum do
        {:ok, contents} = File.read("puzzle1_input.txt")
        masses = contents
            |> String.split
            |> Enum.map(&String.to_integer/1)
        Enum.reduce(masses, 0, &((div(&1, 3) - 2) + &2))
    end
end

IO.puts Puzzle1.calculate_fuel_sum