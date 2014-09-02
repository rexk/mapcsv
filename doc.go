// Copyright 2014 Myung Sun Kim. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Pckage csv reads and writes comma-separated values (CSV) files
//
// Only difference between go's csv package is this package returns
// map[string]string instead of [][]string
// This package can be used to to more advance matching with struct using
// reflect package since result is alraedy captured as map
//
// Performce is not guaranteed
package mapcsv
