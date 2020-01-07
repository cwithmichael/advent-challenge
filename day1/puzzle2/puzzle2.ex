defmodule Puzzle2 do
    defp calculate_fuel_int(tmp, acc) do
        cond do
            tmp > 0 -> calculate_fuel_int(div(tmp, 3) - 2, tmp+acc)
            tmp <= 0 -> acc
        end
    end

    def calculate_fuel_sum(filename) do
        {:ok, contents} = File.read(filename)
        masses = contents
            |> String.split
            |> Enum.map(&String.to_integer/1)
        Enum.reduce(Enum.map(masses, fn(x) -> calculate_fuel_int(x, -x) end), fn(x, acc) -> x + acc end)
    end
end

IO.puts Puzzle2.calculate_fuel_sum("puzzle1_input.txt")