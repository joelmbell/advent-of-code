import Foundation

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

struct Day1Solver {
    typealias FuelConverter = (Int) -> Int

    let converter: FuelConverter
    init(converter: @escaping FuelConverter) {
        self.converter = converter
    }

    func solve(data: [Int]) -> Int {
        return data.reduce(0) { (acc, next) -> Int in
            return acc + self.converter(next)
        }
    }
}

func part1Fuel(to mass: Int) -> Int {
    return Int(Float(mass) / 3.0) - 2
}

func part2Fuel(to mass: Int) -> Int {
    let calc = part1Fuel(to: mass)
    guard calc > 0 else {
        return 0
    }
    return calc + part2Fuel(to: calc)
}

let data = try! FileLoader()
    .loadFile(named: "day1Input")
    .compactMap { Int($0) }

// Part 1
let part1Solver = Day1Solver(converter: part1Fuel(to:))
part1Solver.solve(data: [12]) == 2
part1Solver.solve(data: [14]) == 2
part1Solver.solve(data: [1969]) == 654
part1Solver.solve(data: [100756]) == 33583
part1Solver.solve(data: data) == 3372756
print("Part 1 Solution: \(part1Solver.solve(data: data))")


// Part 2
let part2Solver = Day1Solver(converter: part2Fuel(to:))
part2Solver.solve(data: [12]) == 2
part2Solver.solve(data: [1969]) == 966
part2Solver.solve(data: [100756]) == 50346
part2Solver.solve(data: data) == 5056279
print("Part 2 Solution: \(part2Solver.solve(data: data))")
