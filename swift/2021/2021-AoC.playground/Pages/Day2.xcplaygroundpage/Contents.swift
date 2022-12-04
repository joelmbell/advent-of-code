//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

enum Command {
    case forward(Int)
    case down(Int)
    case up(Int)

    init(val: String) {
        let vals = val.split(separator: " ").map { String($0) }
        let amount = Int(vals[1])!
        switch vals[0] {
        case "forward":
            self = .forward(amount)
        case "down":
            self = .down(amount)
        case "up":
            self = .up(amount)
        default:
            fatalError("Invalid command: \(val)")
        }
    }
}

func solveDay2Pt1(input: [String]) -> Int {
    let input = input
        .filter { !$0.isEmpty }
        .map { Command(val: $0) }

    var hPos: Int = 0
    var dPos: Int = 0

    for cmd in input {
        switch cmd {
        case .forward(let val):
            hPos += val
        case .down(let val):
            dPos += val
        case .up(let val):
            dPos -= val
        }
    }

    return hPos * dPos
}

func solveDay2Pt2(input: [String]) -> Int {
    let input = input
        .filter { !$0.isEmpty }
        .map { Command(val: $0) }

    var hPos: Int = 0
    var dPos: Int = 0
    var aim: Int = 0

    for cmd in input {
        switch cmd {
        case .forward(let val):
            hPos += val
            if aim > 0 {
                dPos += (aim * val)
            }
        case .down(let val):
            aim += val
        case .up(let val):
            aim -= val
        }
    }

    return hPos * dPos
}

Assert(solveDay2Pt1(input: sampleData), 150)
Assert(solveDay2Pt1(input: data), 1580000)
Assert(solveDay2Pt2(input: sampleData), 900)
Assert(solveDay2Pt2(input: data), 1251263225)
