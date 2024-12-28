package LC_2471;

public class Test {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(11);
        root.left = new TreeNode(3);
        root.right = new TreeNode(43);
        root.left.left = new TreeNode(27);
        root.left.right = new TreeNode(29);
//        root.right.left = new TreeNode();
        root.right.right = new TreeNode(17);
        root.left.right.left = new TreeNode(18);

        Solution solution = new Solution();
        System.out.println(solution.minimumOperations(root));
    }
}
