/*
* victi.ms common library
* Copyright (C) 2017 The victi.ms team
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published
* by the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package log

import "github.com/Sirupsen/logrus"

// Logger is the exported single log instance
var Logger *logrus.Logger

// init executes upon import creating a default logger
func init() {
	Logger = New()
	Logger.Info("A new logger has been created")
}

// New creates a new instance a Logger and returns it
// TODO: This should take options
func New() *logrus.Logger {
	logger := logrus.New()
	return logger
}

// SetLogger points Logger to a specific logger instance
func SetLogger(logger *logrus.Logger) {
	Logger = logger
}
