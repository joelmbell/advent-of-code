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

func solveDay1Pt1(data: [String]) -> Int {
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

func solveDay1Pt2(data: [String]) -> Int {
    
    let data = data
        .filter { !$0.isEmpty }
        .map { $0.map { Character("1") == $0 } }
    

    func findRating(in data: [[Bool]], isOxygenRating: Bool) -> [Bool] {
        var data = data
        let columns = data.first?.count ?? 0
        for i in 0..<columns {
            var onCount = 0
            var offCount = 0
            for k in 0..<data.count {
                if data[k][i] {
                    onCount += 1
                } else {
                    offCount += 1
                }
            }
            
            data = data.filter { item in
                if isOxygenRating {
                    return item[i] == (onCount >= offCount)
                } else {
                    return item[i] == (onCount < offCount)
                }
            }

            if data.count <= 1 {
                return data.first!
            }
        }
        
        fatalError("Unable to find solution")
    }
    
    let oxygen = convertToDecimal(findRating(in: data, isOxygenRating: true))
    let co2 = convertToDecimal(findRating(in: data, isOxygenRating: false))
    return oxygen * co2
}

Assert(solveDay1Pt1(data: sampleData), 198)
Assert(solveDay1Pt1(data: data), 4138664)

Assert(solveDay1Pt2(data: sampleData), 230)
Assert(solveDay1Pt2(data: data), 4273224)

//: [Next](@next)
