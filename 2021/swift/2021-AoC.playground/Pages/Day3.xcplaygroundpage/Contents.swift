//: [Previous](@previous)

import Foundation

let sampleData = try! FileLoader().loadFile(named: "sample")
let data = try! FileLoader().loadFile(named: "input")

func convertToDecimal(_ input: [Bool]) -> Int {
    let binaryString = input.reduce(into: "") { partialResult, item in
        let val = item ? 1 : 0
        partialResult += "\(val)"
    }
    
    return Int(binaryString, radix: 2)!
}

func solveDay1(data: [String]) -> Int {
    let data = data
        .filter { !$0.isEmpty }
        .map { $0.map { Character("1") == $0 } }

    var results: [(on: Int, off: Int)] = []
    for i in 0..<data[0].count {
        var countOn = 0
        var countOff = 0
        for k in 0..<data.count {
            if data[k][i] {
                countOn += 1
            } else {
                countOff += 1
            }
        }
        results.append((countOn, countOff))
    }

    var gamma: [Bool] = []
    var epsilon: [Bool] = []
    for result in results {
        if result.on == result.off {
            fatalError("Same results")
        }

        gamma.append(result.on > result.off)
        epsilon.append(result.on < result.off)
    }

    return convertToDecimal(epsilon) * convertToDecimal(gamma)
}

Assert(solveDay1(data: sampleData), 198)
Assert(solveDay1(data: data), 4138664)

//: [Next](@next)
