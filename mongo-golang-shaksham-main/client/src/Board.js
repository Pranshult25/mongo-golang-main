import BoardHeader from "./BoardHeader";
import PostForm from "./PostForm";
import PostsListingByCategory from "./PostListingByCategory";
import PostsListing from "./PostsListing";

function Board() {
  return (<div style={{minHeight:"100vh",backgroundColor:"black"}}>
    <BoardHeader />
    <PostForm />
    <PostsListing />
  </div>);
}

export default Board;