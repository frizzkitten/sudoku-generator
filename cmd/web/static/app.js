document.addEventListener('DOMContentLoaded', () => {
    const generateBtn = document.getElementById('generate-btn');
    const baseInput = document.getElementById('base-input');
    const sudokuContainer = document.getElementById('sudoku-container');

    generateBtn.addEventListener('click', generateSudoku);

    // Generate on Enter key
    baseInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
            generateSudoku();
        }
    });

    // Generate initial puzzle
    generateSudoku();

    async function generateSudoku() {
        const base = parseInt(baseInput.value);

        if (!base || base < 1 || base > 5) {
            showError('Please enter a base between 1 and 5');
            return;
        }

        sudokuContainer.innerHTML = '<div class="loading">Generating puzzle...</div>';

        try {
            const response = await fetch(`/generate?base=${base}`);

            if (!response.ok) {
                throw new Error('Failed to generate sudoku');
            }

            const data = await response.json();
            renderSudoku(data);
        } catch (error) {
            showError('Error generating sudoku: ' + error.message);
        }
    }

    function renderSudoku(data) {
        const { rows, base, size } = data;

        const grid = document.createElement('div');
        grid.className = 'sudoku-grid';
        grid.style.gridTemplateColumns = `repeat(${size}, 45px)`;

        rows.forEach((row, rowIndex) => {
            row.forEach((value, colIndex) => {
                const cell = document.createElement('div');
                cell.className = 'sudoku-cell';
                cell.textContent = value;

                // Add thick border on the right edge of each box (except last column)
                if ((colIndex + 1) % base === 0 && colIndex !== size - 1) {
                    cell.classList.add('border-right-thick');
                }

                // Add thick border on the bottom edge of each box (except last row)
                if ((rowIndex + 1) % base === 0 && rowIndex !== size - 1) {
                    cell.classList.add('border-bottom-thick');
                }

                grid.appendChild(cell);
            });
        });

        sudokuContainer.innerHTML = '';
        sudokuContainer.appendChild(grid);
    }

    function showError(message) {
        sudokuContainer.innerHTML = `<div class="error">${message}</div>`;
    }
});
