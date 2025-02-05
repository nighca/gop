/*
 * Copyright (c) 2021 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package spx

import (
	"github.com/goplus/gop/cl/internal/spx/pkg"
)

type Sprite struct {
	pos pkg.Vector
}

func (p *Sprite) SetCostume(costume interface{}) {
}

func (p *Sprite) Say(msg string, secs ...float64) {
}

func (p *Sprite) Position() *pkg.Vector {
	return &p.pos
}

func Gopt_Sprite_Clone__0(sprite interface{}) {
}

func Gopt_Sprite_Clone__1(sprite interface{}, data interface{}) {
}
