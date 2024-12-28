package LC_2471;

import java.util.LinkedList;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}


public class Solution {
    int count = 0;

    public int minimumOperations(TreeNode root) {
        if (root == null) return 0;

        LinkedList<TreeNode> queue = new LinkedList<>();
        queue.add(root);

        while (!queue.isEmpty()) {
            int level = queue.size();
            Integer[] levelList = new Integer[queue.size()];

            for (int i = 0; i < level; i++) {
                TreeNode current = queue.removeFirst();
                levelList[i] = current.val;

                if(current.left != null) {
                    queue.add(current.left);
                }

                if (current.right != null) {
                    queue.add(current.right);
                }
            }

            sort(levelList);
        }

        return count;
    }

    public void sort(Integer[] levelList) {
        int i = 0;

        while (i < levelList.length - 1) {
            int j = i + 1;
            int minIndex = i;
            while (j < levelList.length) {
                if (levelList[minIndex] > levelList[j]) {
                    minIndex = j;
                }
                j++;
            }
            int temp = levelList[i];
            levelList[i] = levelList[minIndex];
            levelList[minIndex] = temp;

            if (i != minIndex) count += 1;

            i++;
        }
    }
}