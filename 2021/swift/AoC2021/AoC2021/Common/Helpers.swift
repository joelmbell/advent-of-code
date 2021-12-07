import Foundation

public extension Character {
    init?(_ uint16: UInt16) {
        guard let scalar = Unicode.Scalar(uint16) else {
            return nil
        }
        self.init(scalar)
    }
}

struct ParkBenchTimer {

    let startTime: CFAbsoluteTime
    var endTime: CFAbsoluteTime?

    init() {
        startTime = CFAbsoluteTimeGetCurrent()
    }

    mutating func stop() -> CFAbsoluteTime {
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
