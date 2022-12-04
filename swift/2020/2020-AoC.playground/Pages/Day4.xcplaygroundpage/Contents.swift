import Foundation

let data = try! FileLoader()
    .loadFile(named: "input")

struct Passport {
    private var map: [String: String] = [:]
        
    mutating func add(item: String) {
        let data = item.components(separatedBy: ":")
        map[data[0]] = data[1]
    }
    
    var isValidPt1: Bool {
        return  hasRequiredFields
    }
    
    var isValidPt2: Bool {
        return  isByrValid && 
                isIyrValid &&
                isEyrValid &&
                isHgtValid &&
                isHclValid &&
                isEclValid &&
                isPidValid
    }
    
    private var hasRequiredFields: Bool {
        let requiredKeys: [String]  = [
            "byr",
            "iyr",
            "eyr",
            "hgt",
            "hcl",
            "ecl",
            "pid"
//            "cid"
        ]
        
        var missingFields: [String] = []
        for key in requiredKeys {
            let isMissing = map.index(forKey: key) == nil 
            if isMissing {
                missingFields.append(key)
            }
        }
                
        return missingFields.isEmpty
    }
    
    private var isByrValid: Bool {
        guard let value = map["byr"] else {
            return false
        }
        guard value.count == 4 else {
            return false
        }
        let year = Int(value)!
        return  year >= 1920 && year <= 2002
    }
    
    private var isIyrValid: Bool {
        guard let value = map["iyr"] else {
            return false
        }
        
        guard value.count == 4 else {
            return false
        }
        let year = Int(value)!
        return  year >= 2010 && year <= 2020
    }
    
    private var isEyrValid: Bool {
        guard let value = map["eyr"] else {
            return false
        }
        
        guard value.count == 4 else {
            return false
        }
        
        let year = Int(value)!
        return  year >= 2020 && year <= 2030
    }
    
    private var isHgtValid: Bool {        
        guard let value = map["hgt"] else {
            return false
        }
        
        let str = value as NSString
        let amount = Int(str.substring(to: str.length - 2))!
        if value.range(of: "cm") != nil {
            return amount >= 150 && amount <= 193
        }
        
        if value.range(of: "in") != nil {
            return amount >= 59 && amount <= 76
        }
        
        return false
    }
    
    private var isHclValid: Bool {
        guard let value = map["hcl"] else {
            return false
        }
        
        guard value.count == 7 else {
            return false
        }
        
        for (idx, char) in value.enumerated() {
            if idx == 0 {
                guard char == Character("#") else {
                    return false
                }
                continue
            }
            
            guard "0123456789abcdefABCDEF".contains(char) else {
                return false
            }
        }
        
        return true
    }
    
    private var isEclValid: Bool {
        guard let value = map["ecl"] else {
            return false
        }
        let validEyeColors = Set(["amb", "blu", "brn", "gry", "grn", "hzl","oth"])
        return validEyeColors.contains(value)
    }
    
    private var isPidValid: Bool {
        guard let value = map["pid"] else {
            return false
        }
        return Int(value) != nil && value.count == 9
    }
}

var validPassports: Int = 0
var currentPassport = Passport()
for row in data {
    if row.isEmpty {
        let isValid = currentPassport.isValidPt2
        if isValid {
            validPassports += 1
        }
        currentPassport = Passport()
        continue
    }
    
    let comps = row.components(separatedBy: " ")
    for comp in comps {
        currentPassport.add(item: comp)
    }
}

print("valid passport count: \(validPassports)")
