# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        ans = []

        def dfs(node):
            if not node:
                ans.append(None)
                return
            ans.append(node.val)
            dfs(node.left)
            dfs(node.right)

        dfs(root)

        return ','.join([str(x) for x in ans])

        
    # 1,2,None,None,3,None,None
    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
       

        nodes=data.split(',')
        if nodes[0]=='None':
            return None
        
      
        index=0

        def dfs(node):
            nonlocal index
            if nodes[index]=='None':
                index+=1
                return None
            node=TreeNode(int(nodes[index]))
            index+=1
            node.left=dfs(node.left)
            node.right=dfs(node.right)
            return node
        ans=dfs(ans)
        return ans

        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))