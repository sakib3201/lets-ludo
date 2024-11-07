/**
 * Initializes the canvas for the Ludo board.
 * Sets the canvas size to 80% of the smaller dimension of the window
 * and applies a border style.
 *
 * @returns {HTMLCanvasElement} The initialized canvas element.
 */
export function initializeCanvas() {
    const canvas = document.getElementById('ludo_board');
    const { innerWidth: width, innerHeight: height } = window;
    const ludoBoardSize = Math.min(width, height);
    canvas.width = ludoBoardSize * 0.8;
    canvas.height = ludoBoardSize * 0.8;

    canvas.style.border = '5px solid white';
    return canvas;
}

/**
 * Draws the Ludo board on the given canvas.
 *
 * @param {HTMLCanvasElement} canvas
 */
export function drawBoard(canvas) {
    const ctx = canvas.getContext('2d');
    const squareWidth = Math.min(canvas.width / 2, canvas.height / 2);
    const borderGap = squareWidth / 25;
    const laneGap = squareWidth / 4;

    // Top left rectangle
    ctx.fillStyle = 'red';
    ctx.fillRect(borderGap, borderGap, squareWidth - laneGap, squareWidth - laneGap);

    // Top right rectangle
    ctx.fillStyle = 'green';
    ctx.fillRect(squareWidth + laneGap, borderGap, squareWidth - laneGap - borderGap, squareWidth - laneGap);

    // Bottom left rectangle
    ctx.fillStyle = 'blue';
    ctx.fillRect(borderGap, squareWidth + laneGap - borderGap, squareWidth - laneGap, squareWidth - laneGap);

    // Bottom right rectangle
    ctx.fillStyle = 'yellow';
    ctx.fillRect(squareWidth + laneGap, squareWidth + laneGap - borderGap, squareWidth - laneGap - borderGap, squareWidth - laneGap);

    // center rectangle
    ctx.fillStyle = 'white';
    ctx.beginPath();
    ctx.arc(squareWidth + borderGap / 2, squareWidth, laneGap / 2, 0, 2 * Math.PI);
    ctx.fill();
    ctx.closePath();
}
