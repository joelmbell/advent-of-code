//
//  Day8.swift
//  AoC2021
//
//  Created by Bell, Joel on 12/7/21.
//

import Foundation

private struct Decoder {
    
    private let key: [String]
    private let secret: [String]
    
    init(_ line: String) {
        let parts = line.components(separatedBy: " | ")
        key = parts[0].components(separatedBy: " ")
        secret = parts[1].components(separatedBy: " ")    
    }
    
    func solve() -> Int {
        var counter = 0
        for s in secret {
            if s.count == 7 { // is 8
                counter += 1
            } else if s.count == 3 { // is 7
                counter += 1
            } else if s.count == 4 { // is 4
                counter += 1
            } else if s.count == 2 { // is 1
                counter += 1
            }
        }
        return counter
    }
}

private func solvePt1(_ data: [String]) -> Int {
    let decoders = data
        .filter { !$0.isEmpty }
        .map { Decoder($0) }
    
    return decoders.reduce(into: 0) { partialResult, decoder in
        partialResult += decoder.solve()
    }
}

func runDay8() {
    let sampleData = try! FileLoader().loadFile(at: "Day8/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day8/input.txt")
    
    print("pt 1 sample: \(solvePt1(sampleData) == 26)")
    print("pt 1 data: \(solvePt1(data))")

//    print("pt 2 sample: \(solvePt2(sampleData) == 168)")
    
//    var timer = ParkBenchTimer()
//    print("pt 2 data: \(solvePt2(data) == 96864235)")
//    print("\tpt 2 data runtime: \(timer.stop())")
}
