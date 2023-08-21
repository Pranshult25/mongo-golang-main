import BoardHeader from "./BoardHeader";
import PostForm from "./PostForm";
import PostsListing from "./PostsListing";

function Board() {
  return (<div style={{minHeight:"100vh",backgroundColor:"black"}}>
    <BoardHeader />
    <PostForm />
    <PostsListing />
  </div>);
}

export default Board;