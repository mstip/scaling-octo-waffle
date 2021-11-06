import {
    ChangeDetectorRef,
    Component,
    ElementRef,
    EventEmitter,
    Input,
    OnInit,
    Output,
    ViewChild
} from '@angular/core';
import * as d3 from 'd3';

@Component({
    selector: 'app-structure-tree',
    templateUrl: './tree.component.html',
    styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {

    @Input() assets: any[] = [];
    @Output() onElement = new EventEmitter<number>();

    @ViewChild('tree', {static: false, read: ElementRef})
    treeContainer?: ElementRef;

    constructor(private changeDetectorRef: ChangeDetectorRef) {
    }

    ngOnInit(): void {
        this.drawTree();
    }

    drawTree() {
        this.changeDetectorRef.detectChanges();
        const svg = d3.select(this.treeContainer?.nativeElement);
        const treeData = d3.stratify()
            .id(function (d: any) {
                return d.name;
            })
            .parentId(function (d: any) {
                return d.parent;
            })(this.assets);
        const marginLeft = 250;
        const marginTop = 20;

        const dy = 120;
        const dx = 12;

        const tree = d3.tree().nodeSize([50, 50]);
        const treeLink = d3.linkHorizontal().x((d: any) => d.x).y((d: any) => d.y);
        const root = tree(treeData);

        let x0 = Infinity;
        let x1 = -x0;
        root.each(d => {
            if (d.x > x1) x1 = d.x;
            if (d.x < x0) x0 = d.x;
        });

        // @ts-ignore
        svg.attr('viewBox', [0, 0, 500, x1 - x0 + dx * 2])
            .style('overflow', 'visible');

        const g = svg.append('g')
            // .attr('font-family', 'sans-serif')
            .attr('font-size', 4)
            .attr('transform', `translate(${marginLeft},${marginTop})`);


        const link = g.append('g')
            .attr('fill', 'none')
            .attr('stroke', '#a8dadc')
            // .attr('stroke-opacity', 0.4)
            .attr('stroke-width', 0.5)
            .selectAll('path')
            .data(root.links())
            .join('path');

        // @ts-ignore
        link.attr('d', treeLink);

        const node = g.append('g')
            .attr('stroke-linejoin', 'round')
            .attr('stroke-width', 3)
            .selectAll('g')
            .data(root.descendants())
            .join('g')
            .attr('transform', d => `translate(${d.x},${d.y})`)
            .on('click', event => this.onElement.emit(0))
            .on('mouseover', el => el.srcElement.setAttribute('fill', '#a8dadc'))
            .on('mouseout', el => el.srcElement.setAttribute('fill', 'black'));

        node.append('path')
            .attr('d', 'M3 2.5a2.5 2.5 0 0 1 5 0 2.5 2.5 0 0 1 5 0v.006c0 .07 0 .27-.038.494H15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1v7.5a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 1 14.5V7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h2.038A2.968 2.968 0 0 1 3 2.506V2.5zm1.068.5H7v-.5a1.5 1.5 0 1 0-3 0c0 .085.002.274.045.43a.522.522 0 0 0 .023.07zM9 3h2.932a.56.56 0 0 0 .023-.07c.043-.156.045-.345.045-.43a1.5 1.5 0 0 0-3 0V3zM1 4v2h6V4H1zm8 0v2h6V4H9zm5 3H9v8h4.5a.5.5 0 0 0 .5-.5V7zm-7 8V7H2v7.5a.5.5 0 0 0 .5.5H7z')
            .attr('fill', 'black')
            .attr('transform', `scale(0.3)`);

        node.append('text')
            .attr('dy', '1em')
            .attr('x', '1.5em')
            .text((d: any) => d.data.name)
            .clone(true).lower()
            .attr('stroke', 'white');
    }

}
