package objects

import (
	"github.com/wieku/danser-go/app/audio"
	"github.com/wieku/danser-go/app/beatmap/difficulty"
	"github.com/wieku/danser-go/framework/math/vector"
)

type IHitObject interface {
	Update(time int64) bool
	SetTiming(timings *Timings)
	UpdateStacking()
	SetDifficulty(difficulty *difficulty.Difficulty)

	GetStartTime() int64
	GetEndTime() int64

	GetPositionAt(int64) vector.Vector2f
	GetStackedPositionAt(int64) vector.Vector2f
	GetStackedPositionAtMod(time int64, modifier difficulty.Modifier) vector.Vector2f

	GetStartPosition() vector.Vector2f
	GetStackedStartPosition() vector.Vector2f
	GetStackedStartPositionMod(modifier difficulty.Modifier) vector.Vector2f

	GetEndPosition() vector.Vector2f
	GetStackedEndPosition() vector.Vector2f
	GetStackedEndPositionMod(modifier difficulty.Modifier) vector.Vector2f

	GetID() int64
	SetID(int64)
	SetComboNumber(cn int64)
	SetComboSet(set int64)

	GetStackIndex(modifier difficulty.Modifier) int64
	SetStackIndex(index int64, modifier difficulty.Modifier)
	SetStackOffset(offset float32, modifier difficulty.Modifier)

	IsNewCombo() bool
}

type HitObject struct {
	StartPosRaw  vector.Vector2f
	EndPosRaw    vector.Vector2f

	StartTime int64
	EndTime   int64

	StackOffset   vector.Vector2f
	StackOffsetEZ vector.Vector2f
	StackOffsetHR vector.Vector2f

	PositionDelegate func(time int64) vector.Vector2f

	StackIndex   int64
	StackIndexEZ int64
	StackIndexHR int64

	HitObjectID int64

	NewCombo    bool
	ComboNumber int64
	ComboSet    int64

	BasicHitSound audio.HitSoundInfo
}

func (hitObject *HitObject) Update(time int64) bool { return true }
func (hitObject *HitObject) SetTiming(timings *Timings) {}
func (hitObject *HitObject) UpdateStacking() {}
func (hitObject *HitObject) SetDifficulty(difficulty *difficulty.Difficulty) {}

func (hitObject *HitObject) GetStartTime() int64 {
	return hitObject.StartTime
}

func (hitObject *HitObject) GetEndTime() int64 {
	return hitObject.EndTime
}

func (hitObject *HitObject) GetPositionAt(time int64) vector.Vector2f {
	if hitObject.PositionDelegate != nil {
		return hitObject.PositionDelegate(time)
	}

	return hitObject.StartPosRaw
}

func (hitObject *HitObject) GetStackedPositionAt(time int64) vector.Vector2f {
	return hitObject.GetPositionAt(time).Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedPositionAtMod(time int64, modifier difficulty.Modifier) vector.Vector2f {
	basePosition := hitObject.GetPositionAt(time)

	switch {
	case modifier&difficulty.HardRock > 0:
		basePosition.Y = 384 - basePosition.Y
		return basePosition.Add(hitObject.StackOffsetHR)
	case modifier&difficulty.Easy > 0:
		return basePosition.Add(hitObject.StackOffsetEZ)
	}

	return basePosition.Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStartPosition() vector.Vector2f {
	return hitObject.StartPosRaw
}

func (hitObject *HitObject) GetStackedStartPosition() vector.Vector2f {
	return hitObject.GetStartPosition().Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedStartPositionMod(modifier difficulty.Modifier) vector.Vector2f {
	basePosition := hitObject.GetStartPosition()

	switch {
	case modifier&difficulty.HardRock > 0:
		basePosition.Y = 384 - basePosition.Y
		return basePosition.Add(hitObject.StackOffsetHR)
	case modifier&difficulty.Easy > 0:
		return basePosition.Add(hitObject.StackOffsetEZ)
	}

	return basePosition.Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetEndPosition() vector.Vector2f {
	return hitObject.EndPosRaw
}

func (hitObject *HitObject) GetStackedEndPosition() vector.Vector2f {
	return hitObject.GetEndPosition().Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedEndPositionMod(modifier difficulty.Modifier) vector.Vector2f {
	basePosition := hitObject.GetEndPosition()

	switch {
	case modifier&difficulty.HardRock > 0:
		basePosition.Y = 384 - basePosition.Y
		return basePosition.Add(hitObject.StackOffsetHR)
	case modifier&difficulty.Easy > 0:
		return basePosition.Add(hitObject.StackOffsetEZ)
	}

	return basePosition.Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetID() int64 {
	return hitObject.HitObjectID
}

func (hitObject *HitObject) SetID(id int64) {
	hitObject.HitObjectID = id
}

func (hitObject *HitObject) SetComboNumber(cn int64) {
	hitObject.ComboNumber = cn
}

func (hitObject *HitObject) SetComboSet(set int64) {
	hitObject.ComboSet = set
}

func (hitObject *HitObject) GetStackIndex(modifier difficulty.Modifier) int64 {
	switch {
	case modifier&difficulty.HardRock > 0:
		return hitObject.StackIndexHR
	case modifier&difficulty.Easy > 0:
		return hitObject.StackIndexEZ
	default:
		return hitObject.StackIndex
	}
}

func (hitObject *HitObject) SetStackIndex(index int64, modifier difficulty.Modifier) {
	switch {
	case modifier&difficulty.HardRock > 0:
		hitObject.StackIndexHR = index
	case modifier&difficulty.Easy > 0:
		hitObject.StackIndexEZ = index
	default:
		hitObject.StackIndex = index
	}
}

func (hitObject *HitObject) SetStackOffset(offset float32, modifier difficulty.Modifier) {
	switch {
	case modifier&difficulty.HardRock > 0:
		hitObject.StackOffsetHR = vector.NewVec2f(1, 1).Scl(offset)
	case modifier&difficulty.Easy > 0:
		hitObject.StackOffsetEZ = vector.NewVec2f(1, 1).Scl(offset)
	default:
		hitObject.StackOffset = vector.NewVec2f(1, 1).Scl(offset)
	}
}

func (hitObject *HitObject) IsNewCombo() bool {
	return hitObject.NewCombo
}

func ModifyPosition(hitObject *HitObject, basePosition vector.Vector2f, modifier difficulty.Modifier) vector.Vector2f {
	switch {
	case modifier&difficulty.HardRock > 0:
		basePosition.Y = 384 - basePosition.Y
		return basePosition.Add(hitObject.StackOffsetHR)
	case modifier&difficulty.Easy > 0:
		return basePosition.Add(hitObject.StackOffsetEZ)
	}

	return basePosition.Add(hitObject.StackOffset)
}