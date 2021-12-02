//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

struct Command {
    enum Direction: String {
        case forward
        case down
        case up
    }

    let direction: Direction
    let amount: Int

    init(val: String) {
        let vals = val.split(separator: " ").map { String($0) }
        direction = Direction(rawValue: vals[0])!
        amount = Int(vals[1])!
    }
}

func solveDay2Pt1(input: [String]) -> Int {
    let input = input
        .filter { !$0.isEmpty }
        .map { Command(val: $0) }

    var hPos: Int = 0
    var dPos: Int = 0

    for cmd in input {
        switch cmd.direction {
        case .forward:
            hPos += cmd.amount
        case .down:
            dPos += cmd.amount
        case .up:
            dPos -= cmd.amount
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
        switch cmd.direction {
        case .forward:
            hPos += cmd.amount
            if aim > 0 {
                dPos += (aim * cmd.amount)
            }
        case .down:
            aim += cmd.amount
        case .up:
            aim -= cmd.amount
        }
    }

    return hPos * dPos
}

Assert(solveDay2Pt1(input: sampleData), 150)
Assert(solveDay2Pt1(input: data), 1580000)
Assert(solveDay2Pt2(input: sampleData), 900)
Assert(solveDay2Pt2(input: data), 1251263225)
