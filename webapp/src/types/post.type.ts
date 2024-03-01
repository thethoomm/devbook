export type PostType = {
  id: number;
  title: string;
  content: string;
  authorId: number;
  authorUsername: string;
  likes: number;
  createdAt: Date;
}