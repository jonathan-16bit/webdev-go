# Fixed Path Patterns
They do not end with trailing slashes, so they must be **exactly** matched.  

# Subtree Patterns
They end with trailing slashes.  
They are matched whenever the **start** of a request URL path matches this subtree path.  
Which is why `"/"` acts like a catch-all.  

# Conclusion
The most specific pattern wins.  
