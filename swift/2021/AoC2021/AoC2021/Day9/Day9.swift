//
//  Day9.swift
//  AoC2021
//
//  Created by Joel Bell on 12/8/21.
//

import Foundation

private func solvePt1(_ data: [String]) -> Int {

    let heightMap = data
        .filter { !$0.isEmpty }
        .map { $0.map { Int(String($0))! } }

    var riskLevel: Int = 0
    for i in 0..<heightMap.count {
        for k in 0..<heightMap[i].count {
            let val = heightMap[i][k]

            let left = (k > 0) ? heightMap[i][k-1] : val + 1
            let right = (k < heightMap[i].count-1) ? heightMap[i][k+1] : val + 1
            let top = (i > 0) ? heightMap[i-1][k] : val + 1
            let bottom = (i < heightMap.count-1) ? heightMap[i+1][k] : val + 1

            if val < left && val < right && val < top && val < bottom {
                riskLevel += val + 1
            }
        }
    }

    return riskLevel
}

func runDay9() {
    let sampleData = try! FileLoader().loadFile(at: "Day9/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day9/input.txt")

//    print("pt 1 sample: \(solvePt1(sampleData))")
    print("pt 1 data: \(solvePt1(data))")

//    print("pt 2 sample: \(solvePt2(sampleData) == 61229)")
//    print("pt 2 data: \(solvePt2(data) == 968175)")
}
