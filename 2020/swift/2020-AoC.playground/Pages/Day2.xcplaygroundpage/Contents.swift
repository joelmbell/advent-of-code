import Foundation

struct PasswordPolicy {
    let lhs: Int
    let rhs: Int
    let char: Character
    
    init(data: String) {
        let comps = data.components(separatedBy: " ")
        let rangeComps = comps[0].components(separatedBy: "-")
        
        self.lhs = Int(rangeComps[0])!
        self.rhs = Int(rangeComps[1])!
        self.char = Character(comps[1])      
    }
}



struct PasswordData {
    let policy: PasswordPolicy
    let password: String
    
    init?(data: String? = nil) {
        guard let data = data, !data.isEmpty else {
            return nil
        }
        
        let components = data.components(separatedBy: ":")
        self.policy = PasswordPolicy(data: components[0])
        self.password = components[1].trimmingCharacters(in: .whitespacesAndNewlines)
    }
    
    var isValidV1: Bool {
        let charCount = password
            .filter({ policy.char == $0 })
            .count
        
        return charCount >= policy.lhs && charCount <= policy.rhs 
    }
    
    var isValidV2: Bool {
        let pass = password as NSString
        let lhsChar = Character(pass.character(at: policy.lhs - 1))!
        let rhsChar = Character(pass.character(at: policy.rhs - 1))!
        
        if lhsChar == policy.char, rhsChar != policy.char {
            return true
        } else if lhsChar != policy.char, rhsChar == policy.char {
            return true
        } else {
            return false
        }
    }
}

let answer = try! FileLoader()
    .loadFile(named: "input")
    .compactMap { PasswordData(data: $0) }
    .filter { $0.isValidV2 }
    .count


