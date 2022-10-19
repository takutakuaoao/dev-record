# remarkDirective

## 型

```typescript
remarkDirective():
  | void
  | import('unified').Transformer<import('mdast').Root, import('mdast').Root>

export interface Root extends Parent {
    type: 'root';
}

export interface Parent extends UnistParent {
    children: Content[];
}

// UnistParent
export interface Parent<ChildNode extends Node<object> = Node, TData extends object = NodeData<ChildNode>>
    extends Node<TData> {
    /**
     * List representing the children of a node.
     */
    children: ChildNode[];
}

export type NodeData<TNode extends Node<object>> = TNode extends Node<infer TData> ? TData : never;

export interface Node<TData extends object = Data> {
    /**
     * The variant of a node.
     */
    type: string;

    /**
     * Information from the ecosystem.
     */
    data?: TData | undefined;

    /**
     * Location of a node in a source document.
     * Must not be present if a node is generated.
     */
    position?: Position | undefined;
}

export interface Data {
    [key: string]: unknown;
}

export interface Position {
    /**
     * Place of the first character of the parsed source region.
     */
    start: Point;

    /**
     * Place of the first character after the parsed source region.
     */
    end: Point;

    /**
     * Start column at each index (plus start line) in the source region,
     * for elements that span multiple lines.
     */
    indent?: number[] | undefined;
}

export interface Point {
    /**
     * Line in a source file (1-indexed integer).
     */
    line: number;

    /**
     * Column in a source file (1-indexed integer).
     */
    column: number;
    /**
     * Character in a source file (0-indexed integer).
     */
    offset?: number | undefined;
}

export type Content = TopLevelContent | ListContent | TableContent | RowContent | PhrasingContent;
```

Rootの中身

```typescript
Root {
    type: 'root';
    children: (TopLevelContent | ListContent | TableContent | RowContent | PhrasingContent)[];
    data: NodeData<Node>>;
}
```

## Transformer

```ts
export type Transformer<
  Input extends Node = Node,
  Output extends Node = Input
> = (
  node: Input,
  file: VFile,
  next: TransformCallback<Output>
) => Promise<Output | undefined | void> | Output | Error | undefined | void

export type TransformCallback<Tree extends Node = Node> = (
  error?: Error | null | undefined,
  node?: Tree | undefined,
  file?: VFile | undefined
) => void
```
