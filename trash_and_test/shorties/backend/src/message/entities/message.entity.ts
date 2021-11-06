import {Column, Entity, PrimaryGeneratedColumn} from "typeorm";

@Entity()
export class Message {

    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    title: string;

    @Column()
    text: string;

}
