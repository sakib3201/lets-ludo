export function initializeCanvas() {
    const canvas = document.getElementById('ludo_board');
    const { innerWidth: width, innerHeight: height } = window;
    const ludoBoardSize = Math.min(width, height);
    canvas.width = ludoBoardSize * 0.8;
    canvas.height = ludoBoardSize * 0.8;

    canvas.style.border = '1px solid white';
    return canvas;
}

export function drawBoard(canvas) {
    const ctx = canvas.getContext('2d');
    const rectWidth = canvas.width / 2;
    const rectHeight = canvas.height / 2;

    // Top left rectangle
    ctx.fillStyle = 'red';
    ctx.fillRect(10, 10, rectWidth - 30, rectHeight - 30);

    // Top right rectangle
    ctx.fillStyle = 'green';
    ctx.fillRect(rectWidth + 30, 10, rectWidth - 40, rectHeight - 30);

    // Bottom left rectangle
    ctx.fillStyle = 'blue';
    ctx.fillRect(10, rectHeight + 10, rectWidth - 20, rectHeight - 20);

    // Bottom right rectangle
    ctx.fillStyle = 'yellow';
    ctx.fillRect(rectWidth + 10, rectHeight + 10, rectWidth - 20, rectHeight - 20);
}
