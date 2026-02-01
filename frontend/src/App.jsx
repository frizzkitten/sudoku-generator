import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [base, setBase] = useState(3)
  const [puzzleData, setPuzzleData] = useState(null)
  const [userInputs, setUserInputs] = useState({})
  const [validation, setValidation] = useState({})
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [checkMessage, setCheckMessage] = useState(null)

  const generateSudoku = async () => {
    if (base < 1 || base > 5) {
      setError('Please enter a base between 1 and 5')
      return
    }

    setLoading(true)
    setError(null)
    setCheckMessage(null)
    setValidation({})
    setUserInputs({})

    try {
      const response = await fetch(`/generate?base=${base}`)

      if (!response.ok) {
        throw new Error('Failed to generate sudoku')
      }

      const data = await response.json()
      setPuzzleData(data)
    } catch (err) {
      setError('Error generating sudoku: ' + err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleCellChange = (row, col, value) => {
    const key = `${row}-${col}`
    if (value === '' || value === null) {
      const newInputs = { ...userInputs }
      delete newInputs[key]
      setUserInputs(newInputs)
    } else {
      setUserInputs({ ...userInputs, [key]: parseInt(value) })
    }
    // Clear validation when user changes input
    if (validation[key]) {
      const newValidation = { ...validation }
      delete newValidation[key]
      setValidation(newValidation)
    }
    setCheckMessage(null)
  }

  const checkAnswers = () => {
    if (!puzzleData) return

    const { puzzle, solution, size } = puzzleData
    const newValidation = {}
    let allCorrect = true
    let hasFilledAll = true

    for (let row = 0; row < size; row++) {
      for (let col = 0; col < size; col++) {
        if (puzzle[row][col] === 0) {
          const key = `${row}-${col}`
          const userValue = userInputs[key]

          if (userValue === undefined) {
            hasFilledAll = false
          } else if (userValue !== solution[row][col]) {
            newValidation[key] = 'incorrect'
            allCorrect = false
          } else {
            newValidation[key] = 'correct'
          }
        }
      }
    }

    setValidation(newValidation)

    if (allCorrect && hasFilledAll) {
      setCheckMessage('Perfect! You solved it correctly! 🎉')
    } else if (allCorrect && !hasFilledAll) {
      setCheckMessage('All entries are correct so far! Keep going!')
    } else {
      setCheckMessage('Some entries are incorrect. Try again!')
    }
  }

  useEffect(() => {
    generateSudoku()
  }, [])

  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      generateSudoku()
    }
  }

  return (
    <div className="container">
      <h1>Sudoku Puzzle Generator</h1>

      <div className="controls">
        <label htmlFor="base-input">Base Size:</label>
        <input
          type="number"
          id="base-input"
          min="1"
          max="5"
          value={base}
          onChange={(e) => setBase(parseInt(e.target.value))}
          onKeyDown={handleKeyDown}
        />
        <button onClick={generateSudoku}>New Puzzle</button>
        <p className="hint">Tip: Base 3 creates a standard 9×9 sudoku</p>
      </div>

      <div id="sudoku-container">
        {loading && <div className="loading">Generating puzzle...</div>}

        {error && <div className="error">{error}</div>}

        {puzzleData && !error && (
          <div className={`sudoku-wrapper ${loading ? 'loading' : 'loaded'}`}>
            <SudokuGrid
              puzzleData={puzzleData}
              userInputs={userInputs}
              validation={validation}
              onCellChange={handleCellChange}
            />
            <div className="check-section">
              <button className="check-button" onClick={checkAnswers}>
                Check
              </button>
              {checkMessage && (
                <div className={`check-message ${checkMessage.includes('Perfect') ? 'success' : checkMessage.includes('correct so far') ? 'partial' : 'warning'}`}>
                  {checkMessage}
                </div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

function SudokuGrid({ puzzleData, userInputs, validation, onCellChange }) {
  const { puzzle, base, size } = puzzleData

  return (
    <div
      className="sudoku-grid"
      style={{ gridTemplateColumns: `repeat(${size}, 45px)` }}
    >
      {puzzle.map((row, rowIndex) => (
        row.map((value, colIndex) => {
          const key = `${rowIndex}-${colIndex}`
          const isGiven = value !== 0
          const userValue = userInputs[key]
          const validationStatus = validation[key]

          const classes = ['sudoku-cell']
          if (isGiven) {
            classes.push('given')
          } else {
            classes.push('editable')
            if (validationStatus === 'correct') {
              classes.push('correct')
            } else if (validationStatus === 'incorrect') {
              classes.push('incorrect')
            }
          }

          // Add thick border on the right edge of each box (except last column)
          if ((colIndex + 1) % base === 0 && colIndex !== size - 1) {
            classes.push('border-right-thick')
          }

          // Add thick border on the bottom edge of each box (except last row)
          if ((rowIndex + 1) % base === 0 && rowIndex !== size - 1) {
            classes.push('border-bottom-thick')
          }

          if (isGiven) {
            return (
              <div key={key} className={classes.join(' ')}>
                {value}
              </div>
            )
          } else {
            return (
              <input
                key={key}
                type="text"
                inputMode="numeric"
                pattern="[0-9]*"
                maxLength="1"
                className={classes.join(' ')}
                value={userValue || ''}
                onChange={(e) => {
                  const val = e.target.value
                  if (val === '' || (parseInt(val) >= 1 && parseInt(val) <= size)) {
                    onCellChange(rowIndex, colIndex, val)
                  }
                }}
              />
            )
          }
        })
      ))}
    </div>
  )
}

export default App
