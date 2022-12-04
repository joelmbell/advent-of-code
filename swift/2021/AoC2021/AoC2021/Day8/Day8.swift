//
//  Day8.swift
//  AoC2021
//
//  Created by Bell, Joel on 12/7/21.
//

import Foundation


private struct Decoder {
    
    let secret: [String]
    
    private(set) var map: [Int: String] = [:]
    
    init(_ line: String) {
        let parts = line.components(separatedBy: " | ")
        let initialKeys = parts[0].components(separatedBy: " ")
            
        secret = parts[1].components(separatedBy: " ")
        
        var keysLeft: [String] = []
        for k in initialKeys {
            switch k.count {
            case 7:
                map[8] = k
            case 3:
                map[7] = k
            case 4:
                map[4] = k
            case 2:
                map[1] = k
            default:
                keysLeft.append(k)
            }
        }
        
        // find top row
        let top = Set(map[7]!).symmetricDifference(map[1]!).first!
        
        // find a 3
        let charsToLookFor = map[1]!.map { $0 } + [top]
        for (idx, key) in keysLeft.enumerated() {
            if key.contains(charsToLookFor[0]) && key.contains(charsToLookFor[1]) && key.contains(charsToLookFor[2]) && key.count == 5 {
                map[3] = key
                keysLeft.remove(at: idx)
                break
            }
        }
        
        // how to find the 5
        // it contains everything from the four that is not also in the 1 and has 5 characters
        for (idx, key) in keysLeft.enumerated() {
            let leftoverFrom4 = Array(Set(map[4]!).symmetricDifference(map[1]!))
            if key.contains(leftoverFrom4[0]) && key.contains(leftoverFrom4[1]) && key.count == 5 {
                map[5] = key
                keysLeft.remove(at: idx)
                break
            }
        }
        
        // The only one left with 5 characters should be 2
        for (idx, key) in keysLeft.enumerated() {
            if key.count == 5 {
                map[2] = key
                keysLeft.remove(at: idx)
                break
            }
        }
        
        // finding 6
        // has 6 characters, contains all of the characters in 5, does not contain all the characters in 1
        for (idx, key) in keysLeft.enumerated() {
            if key.count == 6 {
                if Set(key).intersection(Set(map[5]!)).symmetricDifference(Set(map[5]!)).count != 0 {
                    // does not have all the characters in 5
                    continue
                }
                let leftover = Set(key).symmetricDifference(map[5]!).first!
                if !map[1]!.contains(leftover) {
                    map[6] = key
                    keysLeft.remove(at: idx)
                    break
                }
            }
        }
        
        // finding 9
        for (idx, key) in keysLeft.enumerated() {
            let combine5and4 = Set(map[5]!).union(Set(map[4]!))
            if Set(key).symmetricDifference(combine5and4).count == 0 {
                map[9] = key
                keysLeft.remove(at: idx)
                break
            }
        }
        
        if keysLeft.count == 1 {
            map[0] = keysLeft[0]
            keysLeft.remove(at: 0)
        }
    }
    
    func decodeSecret() -> Int {
        var output = ""
        for num in secret {
            for item in map {
                if num.sorted() == item.value.sorted() {
                    output += "\(item.key)"
                }
            }
        }
        return Int(output)!
    }
}

private func solvePt1(_ data: [String]) -> Int {
    let decoders = data
        .filter { !$0.isEmpty }
        .map { Decoder($0) }
    
    return decoders.reduce(into: 0) { partialResult, decoder in
        var counter = 0
        for s in decoder.secret {
            switch s.count {
            case 7, 3, 4, 2: counter += 1
            default: continue
            }
        }
        partialResult += counter
    }
}

private func solvePt2(_ data: [String]) -> Int {
    return data
        .filter { !$0.isEmpty }
        .map { Decoder($0) }
        .reduce(into: 0) { partialResult, decoder in
            partialResult += decoder.decodeSecret()
        }
}

func runDay8() {
    let joelTest = try! FileLoader().loadFile(at: "Day8/joeltest.txt")
    let sampleData = try! FileLoader().loadFile(at: "Day8/sample.txt")
    let data = try! FileLoader().loadFile(at: "Day8/input.txt")
    
    print("pt 1 sample: \(solvePt1(sampleData) == 26)")
    print("pt 1 data: \(solvePt1(data) == 449)")    
    
    print("pt 2 test: \(solvePt2(joelTest) == 1234)")
    print("pt 2 sample: \(solvePt2(sampleData) == 61229)")
    print("pt 2 data: \(solvePt2(data) == 968175)")
}
