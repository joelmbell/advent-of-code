import Foundation

let data = try! FileLoader()
    .loadFile(named: "input")
    .compactMap { Int($0) }

func solvePt1BruteForce() -> Int {
    for (i, iVal) in data.enumerated() {
        for (k, kVal) in data.enumerated() {
            if i == k {
                continue
            }
            
            let amount = iVal + kVal
            if amount == 2020 {
                return iVal * kVal
            }
        }
    }    
    return -1
}

func solvePt1() -> Int {
    let sorted = data.sorted()
    for (i, iVal) in sorted.enumerated() {
        for (k, kVal) in sorted.enumerated() {
            if i == k {
                continue
            }
            
            let total = iVal + kVal
            if total == 2020 {
                return iVal * kVal
            } else if total > 2020 {
                break
            }
        }
    }
    return -1
}

func solvePt2BruteForce() -> Int {
    for (i, iVal) in data.enumerated() {
        for (k, kVal) in data.enumerated() {
            for (j, jVal) in data.enumerated() {
                if i == k || i == j || k == j {
                    continue
                }
                
                let amount = iVal + kVal + jVal
                if amount == 2020 {
                    return iVal * kVal * jVal
                }
            }
        }
    }
    return -1
}

func solvePt2() -> Int {
    let sorted = data.sorted()
    for (i, iVal) in sorted.enumerated() {
        for (k, kVal) in sorted.enumerated() {
            
            let _amount = iVal + kVal + sorted[0]
            if _amount > 2020 {
                break
            }
            
            for (j, jVal) in sorted.enumerated() {
                if i == k || i == j || k == j {
                    continue
                }
                
                let amount = iVal + kVal + jVal
                if amount == 2020 {
                    return iVal * kVal * jVal
                } else if amount > 2020 {
                    break
                }
            }
        }
    }
    return -1
}

//solvePt1BruteForce() == 1007331     // 7926 rotations
solvePt1() == 1007331               // 649 rotations
//solvePt2BruteForce() == 48914340    // 411423 rotations
solvePt2() == 48914340              // 1371 rotations

