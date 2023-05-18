%% PI
tic
fprintf("%2.60f \n\n", calculate_pi(1e7))
toc

% 3.141592753589781406020620124763809144496917724609375000000000 
% Elapsed time is 0.018748 seconds.

%% Compression
gene = 'AGCT';
for i = 1 : 9
    gene = append(gene, 'AGCT');
end

fprintf("gene          : %s\n", gene)
fprintf("gene length   : %d\n", length(gene))
whos -regexp gene

compressed = CompressedGene(gene);
disp(compressed.bit_string)
disp(compressed.decompress())
