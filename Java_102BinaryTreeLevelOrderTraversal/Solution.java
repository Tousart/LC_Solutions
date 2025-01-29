package LC102_BinaryTreeLevelOrderTraversal;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();

        if (root == null)
            return result;

        int index = 0;

        // Добавление корня
        result.add(new ArrayList<>());
        result.get(index).add(root.val);
        index++;

        // Добавление узлов
        addElem(root.left, root.right, result, index);

        return result;
    }

    public static void addElem(TreeNode left, TreeNode right, List<List<Integer>> result, int index) {
        if (left == null && right == null)
            return;

        if (result.size() == index) {
            result.add(new ArrayList<>());
        }

        if (left != null && right != null) {
            result.get(index).add(left.val);
            result.get(index).add(right.val);
            index++;
            addElem(left.left, left.right, result, index);
            addElem(right.left, right.right, result, index);
        } else if (left != null) {
            result.get(index).add(left.val);
            index++;
            addElem(left.left, left.right, result, index);
        } else {
            result.get(index).add(right.val);
            index++;
            addElem(right.left, right.right, result, index);
        }
    }
}
