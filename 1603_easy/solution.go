// Problem: https://leetcode.com/problems/design-parking-system/

type ParkingSystem struct {
	bigSlot    int
	mediumSlot int
	smallSlot  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{bigSlot: big, mediumSlot: medium, smallSlot: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if (this.bigSlot - 1) >= 0 {
			this.bigSlot--
			return true
		} else {
			return false
		}
	case 2:
		if (this.mediumSlot - 1) >= 0 {
			this.mediumSlot--
			return true
		} else {
			return false
		}
	case 3:
		if (this.smallSlot - 1) >= 0 {
			this.smallSlot--
			return true
		} else {
			return false
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

//A better solution can be created using maps