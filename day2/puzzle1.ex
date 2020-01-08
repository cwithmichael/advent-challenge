defmodule Puzzle1 do
    def opcode_parser(work_array, idx, max) do
        x = elem(work_array, idx)
        cond do
            x == 99 -> work_array
            idx >= max -> work_array
            idx < max  ->
                y = elem(work_array, elem(work_array, idx+1))
                z = elem(work_array, elem(work_array, idx+2))
                result_idx = elem(work_array, idx+3)
                cond do
                    x == 1 ->
                        work_array = Tuple.delete_at(work_array, result_idx)
                        opcode_parser(Tuple.insert_at(work_array, result_idx, y + z), idx+4, max)
                    x == 2 ->
                        work_array = Tuple.delete_at(work_array, result_idx)
                        opcode_parser(Tuple.insert_at(work_array, result_idx, y * z), idx+4, max)
                    true -> opcode_parser(work_array, idx+1, max)
                end
        end
    end

    def int_cpu(xs) do
        opcode_parser(List.to_tuple(xs), 0, length(xs))
    end

    def compute(filename) do
        {:ok, contents} = File.read(filename)
        ints = contents
            |> String.split(",")
            |> Enum.map(&String.trim_trailing/1)
            |> Enum.map(&String.to_integer/1)
        int_cpu(ints)
    end
end

IO.inspect Puzzle1.compute("input")