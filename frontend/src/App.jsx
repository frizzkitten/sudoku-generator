import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [base, setBase] = useState(3)
  const [sudoku, setSudoku] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const generateSudoku = async () => {
    if (base < 1 || base > 5) {
      setError('Please enter a base between 1 and 5')
      return
    }

    setLoading(true)
    setError(null)

    try {
      const response = await fetch(`/generate?base=${base}`)

      if (!response.ok) {
        throw new Error('Failed to generate sudoku')
      }

      const data = await response.json()
      setSudoku(data)
    } catch (err) {
      setError('Error generating sudoku: ' + err.message)
    } finally {
      setLoading(false)
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
        <button onClick={generateSudoku}>Generate Puzzle</button>
        <p className="hint">Tip: Base 3 creates a standard 9×9 sudoku</p>
      </div>

      {loading && <div className="loading">Generating puzzle...</div>}

      {error && <div className="error">{error}</div>}

      {sudoku && !loading && !error && (
        <div id="sudoku-container">
          <SudokuGrid sudoku={sudoku} />
        </div>
      )}
    </div>
  )
}

function SudokuGrid({ sudoku }) {
  const { rows, base, size } = sudoku

  return (
    <div
      className="sudoku-grid"
      style={{ gridTemplateColumns: `repeat(${size}, 45px)` }}
    >
      {rows.map((row, rowIndex) => (
        row.map((value, colIndex) => {
          const classes = ['sudoku-cell']

          // Add thick border on the right edge of each box (except last column)
          if ((colIndex + 1) % base === 0 && colIndex !== size - 1) {
            classes.push('border-right-thick')
          }

          // Add thick border on the bottom edge of each box (except last row)
          if ((rowIndex + 1) % base === 0 && rowIndex !== size - 1) {
            classes.push('border-bottom-thick')
          }

          return (
            <div key={`${rowIndex}-${colIndex}`} className={classes.join(' ')}>
              {value}
            </div>
          )
        })
      ))}
    </div>
  )
}

export default App
