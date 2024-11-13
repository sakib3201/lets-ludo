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
    const miniBlockWidth = laneGap / 2;
    
    const centerX = (canvas.width + borderGap) /2;
    const centerY = (canvas.height + borderGap) /2;

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

    // Center red triangle
    ctx.beginPath();
    ctx.moveTo(squareWidth + borderGap - laneGap, squareWidth + borderGap - laneGap);
    ctx.lineTo(centerX , centerY);
    ctx.lineTo(squareWidth + borderGap - laneGap , squareWidth - borderGap + laneGap);
    ctx.closePath();
    ctx.fillStyle = 'red';
    ctx.fill();

    // Center green triangle
    ctx.beginPath();
    ctx.moveTo(squareWidth + borderGap - laneGap, squareWidth + borderGap - laneGap);
    ctx.lineTo(centerX , centerY);
    ctx.lineTo(squareWidth + laneGap , squareWidth + borderGap - laneGap);
    ctx.closePath();
    ctx.fillStyle = 'green';
    ctx.fill();

    // Center blue triangle
    ctx.beginPath();
    ctx.lineTo(squareWidth + borderGap - laneGap , squareWidth - borderGap + laneGap);
    ctx.lineTo(centerX , centerY);
    ctx.lineTo(squareWidth + laneGap , squareWidth - borderGap + laneGap);
    ctx.closePath();
    ctx.fillStyle = 'blue';
    ctx.fill();

    // Center yellow triangle
    ctx.beginPath();
    ctx.lineTo(squareWidth + laneGap , squareWidth - borderGap + laneGap);
    ctx.lineTo(centerX, centerY);
    ctx.lineTo(squareWidth + laneGap , squareWidth + borderGap - laneGap);
    ctx.closePath();
    ctx.fillStyle = 'yellow';
    ctx.fill();

    // top and bottom lanes between red and green
    for (let i = 0; i < 6; i++) {
        ctx.fillStyle = 'silver';
        ctx.strokeStyle = 'darkred';
        ctx.lineWidth = 1;
        ctx.fillRect(borderGap + miniBlockWidth * i , squareWidth - laneGap + borderGap, miniBlockWidth, miniBlockWidth);
        ctx.strokeRect(borderGap + miniBlockWidth * i , squareWidth - laneGap + borderGap, miniBlockWidth, miniBlockWidth);
        
        ctx.strokeStyle = 'darkblue';
        ctx.lineWidth = 1;
        ctx.fillRect(borderGap + miniBlockWidth * i , squareWidth - laneGap + 2 * borderGap + 2 * miniBlockWidth, miniBlockWidth, miniBlockWidth);
        ctx.strokeRect(borderGap + miniBlockWidth * i , squareWidth - laneGap + 2 * borderGap + 2 * miniBlockWidth, miniBlockWidth, miniBlockWidth);
        
        ctx.fillStyle = 'red';
        ctx.strokeStyle = 'darkred';
        ctx.lineWidth = 1;
        ctx.fillRect(borderGap + miniBlockWidth * i , squareWidth - laneGap +1.5*borderGap + miniBlockWidth, miniBlockWidth, miniBlockWidth, miniBlockWidth);
        ctx.strokeRect(borderGap + miniBlockWidth * i , squareWidth - laneGap +1.5*borderGap + miniBlockWidth, miniBlockWidth, miniBlockWidth, miniBlockWidth);
    }
}
