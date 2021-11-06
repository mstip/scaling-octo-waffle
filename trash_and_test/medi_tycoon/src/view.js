const gameView = `
<div class="row r0">
<div class="status-field r0 f0"></div>
<div class="status-field r0 f1"></div>
<div class="status-field r0 f2"></div>
<div class="status-field r0 f3"></div>
<div class="status-field r0 f4"></div>
<div class="status-field r0 f5"></div>
</div>
<div class="row r1">
<div class="field r1 f0"></div>
<div class="field r1 f1"></div>
<div class="field r1 f2"></div>
<div class="field r1 f3"></div>
<div class="field r1 f4"></div>
<div class="field r1 f5"></div>
</div>
<div class="row r2">
<div class="field r2 f0"></div>
<div class="field r2 f1"></div>
<div class="field r2 f2"></div>
<div class="field r2 f3"></div>
<div class="field r2 f4"></div>
<div class="field r2 f5"></div>
</div>
<div class="row r3">
<div class="field r3 f0"></div>
<div class="field r3 f1"></div>
<div class="field r3 f2"></div>
<div class="field r3 f3"></div>
<div class="field r3 f4"></div>
<div class="field r3 f5"></div>
</div>
<div class="row r4">
<div class="field r4 f0"></div>
<div class="field r4 f1"></div>
<div class="field r4 f2"></div>
<div class="field r4 f3"></div>
<div class="field r4 f4"></div>
<div class="field r4 f5"></div>
</div>
`;


export default function view($el, workshop, time) {
    $el.innerHTML = gameView;
    $el.querySelector('.r0.f0').innerText = workshop.type;
    $el.querySelector('.r0.f1').innerText = `${workshop.gold} Gold`;
    $el.querySelector('.r0.f2').innerText = `${workshop.worker} Worker`;
    $el.querySelector('.r0.f3').innerText = time.dateText();
    $el.querySelector('.r0.f4').innerText = time.timeText();
    $el.querySelector('.r0.f5').innerText = 'Menu';

    let row = 1;
    let field = 0;
    for (const item of workshop.inventory) {
        $el.querySelector(`.r${row}.f${field}`).innerText = item;
        field += 1;
        if (field > 5) {
            row++;
            field = 0;
        }
    }

}