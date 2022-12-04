import Foundation

public func Assert<T: Equatable>(_ value1: T, _ value2: T) -> Bool {
    guard value1 == value2 else {
        fatalError("\(value1) != \(value2)")
    }
    return true
}

public extension Character {
    init?(_ uint16: UInt16) {
        guard let scalar = Unicode.Scalar(uint16) else {
            return nil
        }
        self.init(scalar)
    }
}

public struct ParkBenchTimer {

    let startTime: CFAbsoluteTime
    var endTime: CFAbsoluteTime?

    public init() {
        startTime = CFAbsoluteTimeGetCurrent()
    }

    public mutating func stop() -> CFAbsoluteTime {
        endTime = CFAbsoluteTimeGetCurrent()
        return duration!
    }

    var duration: CFAbsoluteTime? {
        if let endTime = endTime {
            return endTime - startTime
        } else {
            return nil
        }
    }
}
