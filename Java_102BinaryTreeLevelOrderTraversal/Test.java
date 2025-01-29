package LC102_BinaryTreeLevelOrderTraversal;

import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(3);
        root.left = new TreeNode(9);
        root.right = new TreeNode(20);
        root.left.left = new TreeNode(1);
        root.right.left = new TreeNode(15);
        root.right.right = new TreeNode(7);

        Solution sl = new Solution();
        System.out.println(sl.levelOrder(root));
    }
}
