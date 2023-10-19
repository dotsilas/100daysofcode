const canvas = document.getElementById("canvas1");
const ctx = canvas.getContext("2d");
console.log(ctx);

window.addEventListener("resize", () => {
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
})

const mousePosition = {
  x: undefined,
  y: undefined,
}

canvas.addEventListener("mousemove", (event) => {
  mousePosition.x = event.x;
  mousePosition.y = event.y;
  drawCircle();
})

function drawCircle(){
  ctx.fillStyle = "red";
  ctx.beginPath();
  ctx.arc(mousePosition.x, mousePosition.y, 3, 0, Math.PI * 2);
  ctx.fill();
}
