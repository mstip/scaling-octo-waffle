export default function TodoItem(props: any) {

    console.log(props);
    return (<li>
        {props.done ?
            <s>{props.text}</s> :
            <span>{props.text}</span>
        }
        <input type="checkbox" checked={props.done} onChange={() => props.onToggleDone(props.index)} />
        {props.done ? <button onClick={() => props.onDeleteTodo(props.index)}>Delete</button> : ''}

    </li>)
}