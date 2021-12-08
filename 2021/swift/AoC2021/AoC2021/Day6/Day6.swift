//
//  Day6.swift
//  AoC2021
//
//  Created by Bell, Joel on 12/7/21.
//

import Foundation

private func solve(_ data: [String], numOfDays: Int) -> Int {
    
    let schoolOfFish = data
        .filter { !$0.isEmpty }
        .first!
        .components(separatedBy: ",")
        .map { Int($0)! }
    
    var map: [Int: Int] = [
        0: 0,
        1: 0,
        2: 0,
        3: 0,
        4: 0,
        5: 0,
        6: 0,
        7: 0,
        8: 0
    ]
    
    for fish in schoolOfFish {
        map[fish] = map[fish]! + 1
    }
    
    for _ in 0..<numOfDays {
        map = [
            0: map[1]!,
            1: map[2]!,
            2: map[3]!,
            3: map[4]!,
            4: map[5]!,
            5: map[6]!,
            6: map[7]! + map[0]!,
            7: map[8]!,
            8: map[0]!
        ]
    }
    
    return map.reduce(into: 0) { partialResult, item in
        partialResult += item.value
    }
}

func runDay6() {
    let sampleData = try! FileLoader().loadFile(at: "Day6/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day6/input.txt")
    
    print("pt 1 sample: \(solve(sampleData, numOfDays: 80) == 5934)")
    print("pt 1 data: \(solve(data, numOfDays: 80) == 354564)")

    print("pt 2 sample: \(solve(sampleData, numOfDays: 256) == 26984457539)")
    print("pt 2 data: \(solve(data, numOfDays: 256) == 1609058859115)")
}
