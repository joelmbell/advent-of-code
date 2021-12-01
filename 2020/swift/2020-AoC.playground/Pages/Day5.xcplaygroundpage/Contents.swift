import Foundation
import UIKit

let data = try! FileLoader()
    .loadFile(named: "input")
    .filter { !$0.isEmpty }
    .map { $0 as NSString }

struct SeatData {
    let row: CGFloat
    let column: CGFloat
    var seatId: CGFloat {
        return row * 8 + column
    }
    
    var uniqueId: String {
        return "\(row)x\(column)"
    }
    
    static let planeSize = CGSize(width: 8, height: 128)
}

func search(range: Range<Int>, data: String) -> CGFloat {
    var lowerBound = CGFloat(range.lowerBound)
    var upperBound = CGFloat(range.upperBound)
    
    var found: CGFloat = 0
    
    for char in data {
        if lowerBound + 2 >= upperBound {
            if lowerBound + 2 == upperBound {
                found = lowerBound
            } else {
                found = lowerBound + 1
            }
            break
        }
        
        if char == "F" || char == "L" {
            upperBound -= floor((upperBound - lowerBound) / 2)
        } else if char == "B" || char == "R" {
            lowerBound += ceil((upperBound - lowerBound) / 2)
        }
    }
    
    return found
}

func findSeat(for data: NSString) -> SeatData {
    let rowData = data.substring(with: NSMakeRange(0, 7))
    let row = search(range: 0..<Int(SeatData.planeSize.height-1), data: rowData)

    let columnData = data.substring(with: NSMakeRange(7, 3))    
    let column = search(range: 0..<Int(SeatData.planeSize.width-1), data: columnData)
    
    return SeatData(row: row, column: column)
}

func solvePt1() -> Int {
    let answer = data
        .map { findSeat(for: $0) }
        .sorted { (lhs, rhs) -> Bool in
            lhs.seatId > rhs.seatId
        }
        .first
    
    return Int(answer!.seatId)
}

func solvePt2() {
    var map: [String: SeatData] = [:]
    
    for seatData in data {
        let seat = findSeat(for: seatData)
//        print("seat: \(seat.row), \(seat.column)")
        
        if map.keys.contains(seat.uniqueId) {
            print("duplicate: \(seat.uniqueId)")
        } else {
            print("unique: \(seat.uniqueId)")
        }   
        
        map[seat.uniqueId] = seat
    }
    
//    var seatsThatDontExist: [SeatData] = []
//    var foundSeats: [SeatData] = []
//    for row in 0..<Int(SeatData.planeSize.height) {
//        for column in 0..<Int(SeatData.planeSize.width) {
//            let seat = SeatData(row: CGFloat(row), column: CGFloat(column))
//            
//            if !map.keys.contains(seat.uniqueId) {
//                seatsThatDontExist.append(seat)
////                print("seat at: \(seat.row), \(seat.column) doesn't exist")
//            } else {
//                foundSeats.append(seat)
//            }
//        }
//    }
//    
//    print("dont exist: \(seatsThatDontExist.count)")
//    print("found: \(foundSeats.count)")
//    print("total: \(data.count)")
//    
//    seatsThatDontExist.sorted { (lhs, rhs) -> Bool in
//        lhs.seatId < rhs.seatId
//    }
    
    
}

//print(solvePt1() == 871)
solvePt2()





