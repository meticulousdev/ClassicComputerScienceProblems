% getsizeof Return the size of an object in bytes
% getsizeof(variable, whos)
% 
% a = 1;
% b = 'abc';
% c = linspace(1, 100);
% whos
%   Name      Size             Bytes  Class     Attributes
% 
%   a         1x1                  8  double              
%   b         1x3                  6  char                
%   c         1x100              800  double        
% 
% >> fprintf("%d BYTE\n", getsizeof('b', whos));
% 6 BYTE
function ret_size = getsizeof(var_name, ret_whos)
    ret_size = -1;

    for idx = 1 : length(ret_whos)
        if ret_whos(idx).name == var_name
            ret_size = ret_whos(idx).bytes;
        end
    end

    if ret_size == -1
        error("Invalid variable name: %s", var_name);
    end
end