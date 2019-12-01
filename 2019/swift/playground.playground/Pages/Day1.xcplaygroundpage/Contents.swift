//: [Previous](@previous)

import Foundation

/**
 Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

 For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
 For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
 For a mass of 1969, the fuel required is 654.
 For a mass of 100756, the fuel required is 33583.
 */

/**
 What is the sum of the fuel requirements for all of the modules on your spacecraft?
 */

struct FileLoader {
    enum Err: Error {
        case loadError
    }

    func loadFile(named name: String, ext: String = "txt") throws -> [String] {
        guard
            let filePath = Bundle.main.path(forResource: name, ofType: ext),
            let contentData = FileManager.default.contents(atPath: filePath),
            let contents = String(data: contentData, encoding: .utf8)
            else { throw Err.loadError }

        return contents.components(separatedBy: "\n")
    }
}

struct Day1Part1Solution {
    func solve(data: [Int]) -> Int {
        return data.reduce(0) { (acc, next) -> Int in
            return acc + fuel(for: next)
        }
    }

    private func fuel(for mass: Int) -> Int {
        return max(Int(Float(mass) / 3.0) - 2, 0)
    }
}

func solvePart1() {
    let part1Solver = Day1Part1Solution()

    // Tests
    part1Solver.solve(data: [12]) == 2
    part1Solver.solve(data: [14]) == 2
    part1Solver.solve(data: [1969]) == 654
    part1Solver.solve(data: [100756]) == 33583


    let data = try! FileLoader()
        .loadFile(named: "day1Input")
        .compactMap { Int($0) }
    print("Part 1 Solution: \(part1Solver.solve(data: data))")
}



// PART 2

struct Day1Part2Solution {
    func solve(data: [Int]) -> Int {
        return data.reduce(0) { (acc, next) -> Int in
            return acc + fuel(for: next)
        }
    }

    private func fuel(for mass: Int) -> Int {
        let calc = Int(Float(mass) / 3.0) - 2
        let result = max(calc, 0)
        var acc: Int = result
        if result == 0 {
            return acc
        } else {
            acc += fuel(for: result)
        }
        return acc
    }
}

func solvePart2() {
    let part2Solver = Day1Part2Solution()

    // Tests
    part2Solver.solve(data: [12]) == 2
    part2Solver.solve(data: [1969]) == 966
    part2Solver.solve(data: [100756]) == 50346

    let data = try! FileLoader()
        .loadFile(named: "day1Input")
        .compactMap { Int($0) }
    print("Part 2 Solution: \(part2Solver.solve(data: data))")
}

solvePart1()
solvePart2()





